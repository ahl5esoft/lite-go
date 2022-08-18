import { readdir, writeFile } from 'fs/promises';
import { basename, join } from 'path';
import { argv } from 'process';

(async () => {
    const apiDirPath = join(__dirname, '..', '..', 'api');
    const endpointDirPaths = await readdir(apiDirPath);
    const metadataBf = [
        'package api',
        '',
        'import (',
        '   "github.com/ahl5esoft/lite-go/contract"',
    ];
    const project = basename(
        join(apiDirPath, '..')
    );
    const group = basename(
        join(apiDirPath, '..', '..')
    );
    const nameSpace = argv[2] || 'git.feigo.fun';
    for (const r of endpointDirPaths) {
        if (r.includes('.'))
            continue;

        metadataBf.push(
            `   "${nameSpace}/${group}/${project}/api/${r}"`
        );

        const childFiles = await readdir(
            join(apiDirPath, r)
        );
        for (const cr of childFiles) {
            const name = cr.replace('.go', '');
            const structName = name.split('-').map(r => {
                return r[0].toUpperCase() + r.substring(1);
            }).join('');
            metadataBf.push(`   apiFactory.Register("${r}", "${name}", ${r}.${structName}Api{})`);
        }
    }
    metadataBf.splice(
        3 + endpointDirPaths.length,
        0,
        ')',
        '',
        'func Register(apiFactory contract.IApiFactory) {',
    );
    metadataBf.push('}');
    await writeFile(
        join(apiDirPath, 'metadata.go'),
        metadataBf.join('\r')
    );
})();
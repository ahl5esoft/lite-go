import { service, version } from 'lite-ts';

(async () => {
    const ioFactory = new service.FSIOFactory();
    const readmeFile = ioFactory.buildFile(__dirname, '..', '..', 'README.md');
    const packageJSONFile = ioFactory.buildFile(__dirname, '..', '..', 'package.json');
    const packageLockJSONFile = ioFactory.buildFile(__dirname, '..', '..', 'package-lock.json');
    new version.CheckHandler(process.argv[2]).setNext(
        new version.ReadmeHandler(readmeFile, process.argv[2])
    ).setNext(
        new version.JsonFileHandler(packageJSONFile, process.argv[2])
    ).setNext(
        new version.JsonFileHandler(packageLockJSONFile, process.argv[2])
    ).handle();
})();
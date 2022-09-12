package mongosvc

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/global"
	"go.mongodb.org/mongo-driver/bson"
)

// 用户随机种子长度区间
var UserRandSeedLengthRange = make(map[string][2]int)

// 用户随机种子服务
type userRandSeedService struct {
	dbFactory     contract.IDbFactory
	scene, userID string
	lenRange      [2]int
	entries       []global.UserRandSeed
}

func (m *userRandSeedService) Get(uow contract.IUnitOfWork, length int, offset int) (res int, err error) {
	var seed string
	if seed, err = m.getSeed(uow); err != nil {
		return
	}

	if length+offset > len(seed) {
		err = fmt.Errorf("种子已达最大值")
		return
	}

	res, err = strconv.Atoi(seed[offset : length+offset])
	return
}

func (m *userRandSeedService) Use(uow contract.IUnitOfWork, length int) (res int, err error) {
	var seed string
	if seed, err = m.getSeed(uow); err != nil {
		return
	}

	if length > len(seed) {
		err = fmt.Errorf("种子已达最大值")
		return
	}

	if res, err = strconv.Atoi(seed[:length]); err != nil {
		return
	}

	m.entries[0].Seed[m.scene] = m.entries[0].Seed[m.scene][length:]
	m.dbFactory.Db(global.UserRandSeed{}, uow).Save(m.entries[0])
	return
}

func (m *userRandSeedService) getSeed(uow contract.IUnitOfWork) (res string, err error) {
	db := m.dbFactory.Db(global.UserRandSeed{}, uow)
	if m.entries == nil {
		err = db.Query().Where(bson.M{
			"_id": m.userID,
		}).ToArray(&(m.entries))
		if err != nil {
			return
		}

		if len(m.entries) == 0 {
			m.entries = append(m.entries, global.UserRandSeed{
				ID:   m.userID,
				Seed: make(map[string]string),
			})
			db.Add(m.entries[0])
		}
	}

	if _, ok := m.entries[0].Seed[m.scene]; !ok {
		m.entries[0].Seed[m.scene] = ""
	}

	if len(m.entries[0].Seed[m.scene]) < m.lenRange[0] {
		rand.Seed(
			time.Now().UnixNano(),
		)
		for len(m.entries[0].Seed[m.scene]) < m.lenRange[1] {
			m.entries[0].Seed[m.scene] += strconv.FormatInt(
				rand.Int63(),
				10,
			)
		}
		db.Save(m.entries[0])
	}

	res = m.entries[0].Seed[m.scene]
	return
}

// 创建用户随机种子服务
func NewUserRandSeedService(
	dbFactory contract.IDbFactory,
	scene, userID string,
) contract.IUserRandSeedService {
	if _, ok := UserRandSeedLengthRange[scene]; !ok {
		UserRandSeedLengthRange[scene] = [2]int{128, 512}
	}

	return &userRandSeedService{
		dbFactory: dbFactory,
		scene:     scene,
		userID:    userID,
		lenRange:  UserRandSeedLengthRange[scene],
	}
}

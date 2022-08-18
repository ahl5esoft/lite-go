package goredissvc

import (
	"context"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/ahl5esoft/lite-go/service/redissvc"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/go-redis/redis/v8"
)

var redisAdapterMutex sync.Mutex

type redisAdapter struct {
	client         redis.Cmdable
	clusterOptions *redis.ClusterOptions
	ctx            context.Context
	options        *redis.Options
}

func (r *redisAdapter) BitCount(key string, start, end int64) (int64, error) {
	return r.getClient().BitCount(r.ctx, key, &redis.BitCount{
		End:   end,
		Start: start,
	}).Result()
}

func (r *redisAdapter) BitOp(op, destKey string, keys ...string) (bool, error) {
	var res int64
	var err error
	switch op {
	case "and":
		res, err = r.getClient().BitOpAnd(r.ctx, destKey, keys...).Result()
	case "not":
		res, err = r.getClient().BitOpNot(r.ctx, destKey, keys[0]).Result()
	case "or":
		res, err = r.getClient().BitOpOr(r.ctx, destKey, keys...).Result()
	default:
		res, err = r.getClient().BitOpXor(r.ctx, destKey, keys...).Result()
	}

	if err != nil {
		return false, err
	}

	temp := false
	if res == 1 {
		temp = true
	}
	return temp, err
}

func (r *redisAdapter) BitPos(key string, bit bool, start, end int64) (int64, error) {
	v := int64(0)
	if bit {
		v = 1
	}

	return r.getClient().BitPos(r.ctx, key, v, start, end).Result()
}

func (r *redisAdapter) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return r.getClient().BLPop(r.ctx, timeout, keys...).Result()
}

func (r *redisAdapter) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	return r.getClient().BRPop(r.ctx, timeout, keys...).Result()
}

func (m *redisAdapter) Close() error {
	if c, ok := m.getClient().(*redis.Client); ok {
		return c.Close()
	}

	return m.getClient().(*redis.ClusterClient).Close()
}

func (r *redisAdapter) Decr(key string) (int64, error) {
	return r.getClient().Decr(r.ctx, key).Result()
}

func (r *redisAdapter) DecrBy(key string, decrement int64) (int64, error) {
	return r.getClient().DecrBy(r.ctx, key, decrement).Result()
}

func (r *redisAdapter) Del(keys ...string) (int64, error) {
	if len(keys) == 0 {
		return 0, nil
	}

	return r.getClient().Del(r.ctx, keys...).Result()
}

func (r *redisAdapter) Exists(str string) (bool, error) {
	res := r.getClient().Exists(r.ctx, str)
	return res.Val() == 1, res.Err()
}

func (r *redisAdapter) ExpireAt(key string, time time.Time) (bool, error) {
	return r.getClient().ExpireAt(r.ctx, key, time).Result()
}

func (r *redisAdapter) Expires(key string, seconds time.Duration) (bool, error) {
	return r.getClient().Expire(r.ctx, key, seconds).Result()
}

func (r *redisAdapter) GeoAdd(key string, locations ...message.RedisGeoLocation) (int64, error) {
	var temp []*redis.GeoLocation
	underscore.Chain(locations).Map(func(r message.RedisGeoLocation, _ int) *redis.GeoLocation {
		return &redis.GeoLocation{
			Latitude:  r.Latitude,
			Longitude: r.Longitude,
			Name:      r.Member,
		}
	}).Value(&temp)
	return r.getClient().GeoAdd(r.ctx, key, temp...).Result()
}

func (r *redisAdapter) GeoDist(key string, member1, member2, unit string) (float64, error) {
	return r.getClient().GeoDist(r.ctx, key, member1, member2, unit).Result()
}

func (r *redisAdapter) GeoPos(key string, members ...string) ([]*message.RedisGeoPosition, error) {
	res, err := r.getClient().GeoPos(r.ctx, key, members...).Result()
	if err != nil {
		return nil, err
	}

	var temp []*message.RedisGeoPosition
	underscore.Chain(res).Map(func(r *redis.GeoPos, _ int) *message.RedisGeoPosition {
		if r == nil {
			return nil
		}

		return &message.RedisGeoPosition{
			Latitude:  r.Latitude,
			Longitude: r.Longitude,
		}
	}).Value(&temp)
	return temp, nil
}

func (r *redisAdapter) GeoRadius(key string, longitude, latitude float64, query message.RedisGeoRadiusQuery) ([]message.RedisGeoLocation, error) {
	res, err := r.getClient().GeoRadius(r.ctx, key, longitude, latitude, &redis.GeoRadiusQuery{
		Count:       query.Count,
		Radius:      query.Radius,
		Sort:        query.Sort,
		Unit:        query.Unit,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithHash,
	}).Result()
	if err != nil {
		return nil, err
	}

	var temp []message.RedisGeoLocation
	underscore.Chain(res).Map(func(r redis.GeoLocation, _ int) message.RedisGeoLocation {
		return message.RedisGeoLocation{
			RedisGeoPosition: message.RedisGeoPosition{
				Latitude:  r.Latitude,
				Longitude: r.Longitude,
			},
			Distance: r.Dist,
			Hash:     r.GeoHash,
			Member:   r.Name,
		}
	}).Value(&temp)
	return temp, nil
}

func (r *redisAdapter) GeoRadiusByMember(key string, member string, query message.RedisGeoRadiusQuery) ([]message.RedisGeoLocation, error) {
	res, err := r.getClient().GeoRadiusByMember(r.ctx, key, member, &redis.GeoRadiusQuery{
		Count:       query.Count,
		Radius:      query.Radius,
		Sort:        query.Sort,
		Unit:        query.Unit,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithHash,
	}).Result()
	if err != nil {
		return nil, err
	}

	var temp []message.RedisGeoLocation
	underscore.Chain(res).Map(func(r redis.GeoLocation, _ int) message.RedisGeoLocation {
		return message.RedisGeoLocation{
			RedisGeoPosition: message.RedisGeoPosition{
				Latitude:  r.Latitude,
				Longitude: r.Longitude,
			},
			Distance: r.Dist,
			Hash:     r.GeoHash,
			Member:   r.Name,
		}
	}).Value(&temp)
	return temp, nil
}

func (r *redisAdapter) Get(str string) (string, error) {
	res, err := r.getClient().Get(r.ctx, str).Result()
	if err != nil && err == redis.Nil {
		return "", nil
	}

	return res, err
}

func (r *redisAdapter) GetBit(key string, offset int64) (bool, error) {
	res, err := r.getClient().GetBit(r.ctx, key, offset).Result()
	if err != nil {
		return false, err
	}

	return res == 1, nil
}

func (r *redisAdapter) HDel(key string, fields ...string) (int64, error) {
	return r.getClient().HDel(r.ctx, key, fields...).Result()
}

func (r *redisAdapter) HExists(key, field string) (bool, error) {
	return r.getClient().HExists(r.ctx, key, field).Result()
}

func (r *redisAdapter) HGet(key, field string) (string, error) {
	res, err := r.getClient().HGet(r.ctx, key, field).Result()
	if err != nil && err == redis.Nil {
		return "", nil
	}

	return res, err
}

func (r *redisAdapter) HGetAll(key string) (map[string]string, error) {
	return r.getClient().HGetAll(r.ctx, key).Result()
}

func (r *redisAdapter) HIncrBy(key, field string, increment int64) (int64, error) {
	return r.getClient().HIncrBy(r.ctx, key, field, increment).Result()
}

func (r *redisAdapter) HIncrByFloat(key, field string, increment float64) (float64, error) {
	return r.getClient().HIncrByFloat(r.ctx, key, field, increment).Result()
}

func (r *redisAdapter) HKeys(key string) ([]string, error) {
	return r.getClient().HKeys(r.ctx, key).Result()
}

func (r *redisAdapter) HLen(key string) (int64, error) {
	return r.getClient().HLen(r.ctx, key).Result()
}

func (r *redisAdapter) HMGet(key string, fields ...string) ([]string, error) {
	res, err := r.getClient().HMGet(r.ctx, key, fields...).Result()
	if err != nil {
		return nil, err
	}

	var temp []string
	underscore.Chain(res).Map(func(r interface{}, _ int) string {
		return r.(string)
	}).Value(&temp)
	return temp, nil
}

func (r *redisAdapter) HMSet(key string, fieldOrValues ...string) error {
	var temp []interface{}
	underscore.Chain(fieldOrValues).Map(func(r string, _ int) interface{} {
		return r
	}).Value(&temp)
	_, err := r.getClient().HMSet(r.ctx, key, temp...).Result()
	return err
}

func (r *redisAdapter) HScan(key string, cursor uint64, match string, count int64) (map[string]string, uint64, error) {
	res, cursor, err := r.getClient().HScan(r.ctx, key, cursor, match, count).Result()
	if err != nil {
		return nil, cursor, err
	}

	temp := make(map[string]string)
	underscore.Chain(res).Group(func(_ string, ri int) int {
		return ri / 2
	}).Values().Object().Value(&temp)
	return temp, cursor, nil
}

func (r *redisAdapter) HSet(key, field, value string) (bool, error) {
	res, err := r.getClient().HSet(r.ctx, key, field, value).Result()
	if err != nil {
		return false, err
	}

	return res == 1, nil
}

func (r *redisAdapter) HSetNX(key, field, value string) (bool, error) {
	return r.getClient().HSetNX(r.ctx, key, field, value).Result()
}

func (r *redisAdapter) HStrLen(key, field string) (int64, error) {
	res, err := r.HGet(key, field)
	if err != nil {
		return 0, err
	}

	return int64(len(res)), nil
}

func (r *redisAdapter) HVals(key string) ([]string, error) {
	return r.getClient().HVals(r.ctx, key).Result()
}

func (r *redisAdapter) Incr(key string) (int64, error) {
	return r.getClient().Incr(r.ctx, key).Result()
}

func (r *redisAdapter) IncrBy(key string, increment int64) (int64, error) {
	return r.getClient().IncrBy(r.ctx, key, increment).Result()
}

func (r *redisAdapter) LIndex(key string, index int64) (string, error) {
	return r.getClient().LIndex(r.ctx, key, index).Result()
}

func (r *redisAdapter) LLen(key string) (int64, error) {
	return r.getClient().LLen(r.ctx, key).Result()
}

func (r *redisAdapter) LPop(key string) (string, error) {
	return r.getClient().LPop(r.ctx, key).Result()
}

func (r *redisAdapter) LPush(key string, values ...string) (int64, error) {
	if len(values) == 0 {
		return 0, nil
	}

	var temp []interface{}
	underscore.Chain(values).Map(func(r string, _ int) interface{} {
		return r
	}).Value(&temp)
	return r.getClient().LPush(r.ctx, key, temp...).Result()
}

func (r *redisAdapter) LPushX(key string, value string) (int64, error) {
	return r.getClient().LPushX(r.ctx, key, value).Result()
}

func (r *redisAdapter) LRange(key string, start, stop int64) ([]string, error) {
	return r.getClient().LRange(r.ctx, key, start, stop).Result()
}

func (r *redisAdapter) LRem(key string, count int64, value string) (int64, error) {
	return r.getClient().LRem(r.ctx, key, count, value).Result()
}

func (r *redisAdapter) LSet(key string, index int64, value string) (bool, error) {
	res, err := r.getClient().LSet(r.ctx, key, index, value).Result()
	if err != nil {
		return false, err
	}

	return res == "OK", nil
}

func (r *redisAdapter) LTrim(key string, start, stop int64) (bool, error) {
	res, err := r.getClient().LTrim(r.ctx, key, start, stop).Result()
	if err != nil {
		return false, err
	}

	return res == "OK", nil
}

func (r *redisAdapter) RPop(key string) (string, error) {
	return r.getClient().RPop(r.ctx, key).Result()
}

func (r *redisAdapter) RPush(key string, values ...string) (int64, error) {
	if len(values) == 0 {
		return 0, nil
	}

	var temp []interface{}
	underscore.Chain(values).Map(func(r string, _ int) interface{} {
		return r
	}).Value(&temp)
	return r.getClient().RPush(r.ctx, key, temp...).Result()
}

func (r *redisAdapter) RPushX(key string, value string) (int64, error) {
	return r.getClient().RPushX(r.ctx, key, value).Result()
}

func (r *redisAdapter) SAdd(key string, members ...string) (int64, error) {
	var temp []interface{}
	underscore.Chain(members).Map(func(r string, _ int) interface{} {
		return r
	}).Value(&temp)
	return r.getClient().SAdd(r.ctx, key, temp...).Result()
}

func (r *redisAdapter) SCard(key string) (int64, error) {
	return r.getClient().SCard(r.ctx, key).Result()
}

func (r *redisAdapter) Set(key, value string, extraArgs ...interface{}) (ok bool, err error) {
	var res string
	if len(extraArgs) == 0 {
		res, err = r.getClient().Set(r.ctx, key, value, 0).Result()
		ok = res == "OK"
	} else if len(extraArgs) == 1 {
		if extraArgs[0] == "nx" {
			ok, err = r.getClient().SetNX(r.ctx, key, value, 0).Result()
		} else if extraArgs[0] == "xx" {
			ok, err = r.getClient().SetXX(r.ctx, key, value, 0).Result()
		} else {
			panic("redis set 参数有误")
		}
	} else if len(extraArgs) == 2 {
		t := reflect.ValueOf(extraArgs[1]).Int()
		var expires time.Duration
		if extraArgs[0] == "ex" {
			expires = time.Duration(t) * time.Second
		} else if extraArgs[0] == "px" {
			expires = time.Duration(t) * time.Millisecond
		} else {
			panic("redis set 参数有误")
		}
		res, err = r.getClient().Set(r.ctx, key, value, expires).Result()
		ok = res == "OK"
	} else if len(extraArgs) == 3 {
		t := reflect.ValueOf(extraArgs[1]).Int()
		var expires time.Duration
		if extraArgs[0] == "ex" {
			expires = time.Duration(t) * time.Second
		} else if extraArgs[0] == "px" {
			expires = time.Duration(t) * time.Millisecond
		} else {
			panic("redis set 参数有误")
		}

		if extraArgs[2] == "nx" {
			ok, err = r.getClient().SetNX(r.ctx, key, value, expires).Result()
		} else if extraArgs[2] == "xx" {
			ok, err = r.getClient().SetXX(r.ctx, key, value, expires).Result()
		} else {
			panic("redis set 参数有误")
		}
	} else {
		panic("redis set 参数过多")
	}
	return
}

func (r *redisAdapter) SetBit(key string, offset int64, value bool) (bool, error) {
	temp := 0
	if value {
		temp = 1
	}
	res, err := r.getClient().SetBit(r.ctx, key, offset, temp).Result()
	if err != nil {
		return false, err
	}

	return res == 1, nil
}

func (r *redisAdapter) SIsMember(key, member string) (bool, error) {
	return r.getClient().SIsMember(r.ctx, key, member).Result()
}

func (r *redisAdapter) SMembers(key string) ([]string, error) {
	return r.getClient().SMembers(r.ctx, key).Result()
}

func (r *redisAdapter) SPop(key string) (string, error) {
	res, err := r.getClient().SPop(r.ctx, key).Result()
	if err != nil && err == redis.Nil {
		return "", nil
	}
	return res, nil
}

func (r *redisAdapter) Time() (time.Time, error) {
	return r.getClient().Time(r.ctx).Result()
}

func (r *redisAdapter) TTL(key string) (time.Duration, error) {
	return r.getClient().TTL(r.ctx, key).Result()
}

func (r *redisAdapter) WithContext(ctx context.Context) reflect.Value {
	return reflect.ValueOf(&redisAdapter{
		client: r.getClient(),
		ctx:    ctx,
	})
}

func (r *redisAdapter) ZAdd(key string, members ...message.RedisZMember) (int64, error) {
	if len(members) == 0 {
		return 0, nil
	}

	var temp []*redis.Z
	underscore.Chain(members).Map(func(r message.RedisZMember, _ int) *redis.Z {
		return &redis.Z{
			Member: r.Member,
			Score:  r.Score,
		}
	}).Value(&temp)
	return r.getClient().ZAdd(r.ctx, key, temp...).Result()
}

func (r *redisAdapter) ZCard(key string) (int64, error) {
	return r.getClient().ZCard(r.ctx, key).Result()
}

func (r *redisAdapter) ZCount(key string, min, max float64) (int64, error) {
	return r.getClient().ZCount(
		r.ctx,
		key,
		strconv.FormatFloat(min, 'E', -1, 64),
		strconv.FormatFloat(max, 'E', -1, 64),
	).Result()
}

func (r *redisAdapter) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return r.getClient().ZIncrBy(r.ctx, key, increment, member).Result()
}

func (r *redisAdapter) ZRange(key string, start, stop int64, withScores bool) ([]message.RedisZMember, error) {
	var members []message.RedisZMember
	if withScores {
		res, err := r.getClient().ZRangeWithScores(r.ctx, key, start, stop).Result()
		if err != nil {
			return nil, err
		} else if len(res) > 0 {
			underscore.Chain(res).Map(func(r redis.Z, _ int) message.RedisZMember {
				return message.RedisZMember{
					Member: r.Member.(string),
					Score:  r.Score,
				}
			}).Value(&members)
		}
	} else {
		res, err := r.getClient().ZRange(r.ctx, key, start, stop).Result()
		if err != nil {
			return nil, err
		} else if len(res) > 0 {
			underscore.Chain(res).Map(func(r string, _ int) message.RedisZMember {
				return message.RedisZMember{
					Member: r,
				}
			}).Value(&members)
		}
	}

	return members, nil
}

func (r *redisAdapter) ZRangeByScore(key string, min, max string, opt message.RedisZRangeByScore) ([]message.RedisZMember, error) {
	var members []message.RedisZMember
	if opt.WithScores {
		res, err := r.getClient().ZRangeByScoreWithScores(r.ctx, key, &redis.ZRangeBy{
			Count:  opt.Count,
			Max:    max,
			Min:    min,
			Offset: opt.Offset,
		}).Result()
		if err != nil {
			return nil, err
		} else if len(res) > 0 {
			underscore.Chain(res).Map(func(r redis.Z, _ int) message.RedisZMember {
				return message.RedisZMember{
					Member: r.Member.(string),
					Score:  r.Score,
				}
			}).Value(&members)
		}
	} else {
		res, err := r.getClient().ZRangeByScore(r.ctx, key, &redis.ZRangeBy{
			Count:  opt.Count,
			Max:    max,
			Min:    min,
			Offset: opt.Offset,
		}).Result()
		if err != nil {
			return nil, err
		} else if len(res) > 0 {
			underscore.Chain(res).Map(func(r string, _ int) message.RedisZMember {
				return message.RedisZMember{
					Member: r,
				}
			}).Value(&members)
		}
	}

	return members, nil
}

func (r *redisAdapter) ZRank(key, member string) (int64, error) {
	res, err := r.getClient().ZRank(r.ctx, key, member).Result()
	if err == redis.Nil {
		res = -1
		err = nil
	}
	return res, err
}

func (r *redisAdapter) ZRem(key string, members ...string) (int64, error) {
	var temp []interface{}
	underscore.Chain(members).Map(func(r string, _ int) interface{} {
		return r
	}).Value(&temp)
	return r.getClient().ZRem(r.ctx, key, temp...).Result()
}

func (r *redisAdapter) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	return r.getClient().ZRemRangeByRank(r.ctx, key, start, stop).Result()
}

func (r *redisAdapter) ZRemRangeByScore(key string, min, max float64) (int64, error) {
	return r.getClient().ZRemRangeByScore(
		r.ctx,
		key,
		strconv.FormatFloat(min, 'E', -1, 64),
		strconv.FormatFloat(max, 'E', -1, 64),
	).Result()
}

func (r *redisAdapter) ZRevRange(key string, start, stop int64, withScores bool) ([]message.RedisZMember, error) {
	var members []message.RedisZMember
	if withScores {
		res, err := r.getClient().ZRevRangeWithScores(r.ctx, key, start, stop).Result()
		if err != nil {
			return nil, err
		} else if len(res) > 0 {
			underscore.Chain(res).Map(func(r redis.Z, _ int) message.RedisZMember {
				return message.RedisZMember{
					Member: r.Member.(string),
					Score:  r.Score,
				}
			}).Value(&members)
		}
	} else {
		res, err := r.getClient().ZRevRange(r.ctx, key, start, stop).Result()
		if err != nil {
			return nil, err
		} else if len(res) > 0 {
			underscore.Chain(res).Map(func(r string, _ int) message.RedisZMember {
				return message.RedisZMember{
					Member: r,
				}
			}).Value(&members)
		}
	}

	return members, nil
}

func (r *redisAdapter) ZRevRangeByScore(key string, min, max string, opt message.RedisZRangeByScore) ([]message.RedisZMember, error) {
	var members []message.RedisZMember
	if opt.WithScores {
		res, err := r.getClient().ZRevRangeByScoreWithScores(r.ctx, key, &redis.ZRangeBy{
			Count:  opt.Count,
			Max:    max,
			Min:    min,
			Offset: opt.Offset,
		}).Result()
		if err != nil {
			return nil, err
		} else if len(res) > 0 {
			underscore.Chain(res).Map(func(r redis.Z, _ int) message.RedisZMember {
				return message.RedisZMember{
					Member: r.Member.(string),
					Score:  r.Score,
				}
			}).Value(&members)
		}
	} else {
		res, err := r.getClient().ZRevRangeByScore(r.ctx, key, &redis.ZRangeBy{
			Count:  opt.Count,
			Max:    max,
			Min:    min,
			Offset: opt.Offset,
		}).Result()
		if err != nil {
			return nil, err
		} else if len(res) > 0 {
			underscore.Chain(res).Map(func(r string, _ int) message.RedisZMember {
				return message.RedisZMember{
					Member: r,
				}
			}).Value(&members)
		}
	}

	return members, nil
}

func (r *redisAdapter) ZRevRank(key, member string) (int64, error) {
	return r.getClient().ZRevRank(r.ctx, key, member).Result()
}

func (r *redisAdapter) ZScan(key string, cursor uint64, match string, count int64) ([]message.RedisZMember, uint64, error) {
	res, cursor, err := r.getClient().ZScan(r.ctx, key, cursor, match, count).Result()
	if err != nil {
		return nil, cursor, err
	}

	var members []message.RedisZMember
	underscore.Chain(res).Group(func(_ string, ri int) int {
		return ri / 2
	}).Map(func(r []string, _ int) message.RedisZMember {
		score, _ := strconv.ParseFloat(r[1], 64)
		return message.RedisZMember{
			Member: r[0],
			Score:  score,
		}
	}).Value(&members)
	return members, cursor, nil
}

func (r *redisAdapter) ZScore(key, member string) (float64, error) {
	res, err := r.getClient().ZScore(r.ctx, key, member).Result()
	if err == redis.Nil {
		return 0, nil
	}
	return res, err
}

func (m *redisAdapter) getClient() redis.Cmdable {
	if m.client == nil {
		redisAdapterMutex.Lock()
		defer redisAdapterMutex.Unlock()

		if m.client == nil {
			if m.options != nil {
				m.client = redis.NewClient(m.options)
			} else {
				m.client = redis.NewClusterClient(m.clusterOptions)
			}
		}
	}
	return m.client
}

func NewRedis(options ...redissvc.RedisOption[*redisAdapter]) contract.IRedis {
	r := new(redisAdapter)
	r.ctx = context.Background()
	for _, cr := range options {
		cr(r)
	}
	return r
}

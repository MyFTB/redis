package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type clienter interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	Exists(keys ...string) *redis.IntCmd
	TxPipeline() redis.Pipeliner
	Del(keys ...string) *redis.IntCmd
	Close() error
}

func newClient(cli *redis.Client) clienter {
	return normalClient{client: cli}
}

type normalClient struct {
	client *redis.Client
}

func (n normalClient) Get(key string) *redis.StringCmd {
	return n.client.Get(context.Background(), key)
}

func (n normalClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return n.client.Set(context.Background(), key, value, expiration)
}

func (n normalClient) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return n.client.Expire(context.Background(), key, expiration)
}

func (n normalClient) Exists(keys ...string) *redis.IntCmd {
	return n.client.Exists(context.Background(), keys...)
}

func (n normalClient) TxPipeline() redis.Pipeliner {
	return n.client.TxPipeline()
}

func (n normalClient) Del(keys ...string) *redis.IntCmd {
	return n.client.Del(context.Background(), keys...)
}

func (n normalClient) Close() error {
	return n.client.Close()
}

func newClusterClient(cli *redis.ClusterClient) clienter {
	return clusterClient{client: cli}
}

type clusterClient struct {
	client *redis.ClusterClient
}

func (n clusterClient) Get(key string) *redis.StringCmd {
	return n.client.Get(context.Background(), key)
}

func (n clusterClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return n.client.Set(context.Background(), key, value, expiration)
}

func (n clusterClient) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return n.client.Expire(context.Background(), key, expiration)
}

func (n clusterClient) Exists(keys ...string) *redis.IntCmd {
	return n.client.Exists(context.Background(), keys...)
}

func (n clusterClient) TxPipeline() redis.Pipeliner {
	return n.client.TxPipeline()
}

func (n clusterClient) Del(keys ...string) *redis.IntCmd {
	return n.client.Del(context.Background(), keys...)
}

func (n clusterClient) Close() error {
	return n.client.Close()
}



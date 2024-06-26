// This file was generated by .tools/generate_commands.py on Sat Mar 14 2015 08:58:32 +0800
package commands

var HelpCommands = [][]string{
	{"APPEND", "key value", "KV"},
	{"AUTH", "password", "Server"},
	{"BITCOUNT", "key [start] [end]", "KV"},
	{"BITOP", "operation destkey key [key ...]", "KV"},
	{"BITPOS", "key bit [start] [end]", "KV"},
	{"BLPOP", "key [key ...] timeout", "List"},
	{"BRPOP", "key [key ...] timeout", "List"},
	{"CONFIG GET", "parameter", "Server"},
	{"CONFIG REWRITE", "-", "Server"},
	{"DECR", "key", "KV"},
	{"DECRBY", "key decrement", "KV"},
	{"DEL", "key [key ...]", "KV"},
	{"DUMP", "key", "KV"},
	{"ECHO", "message", "Server"},
	{"EVAL", "script numkeys key [key ...] arg [arg ...]", "Script"},
	{"EVALSHA", "sha1 numkeys key [key ...] arg [arg ...]", "Script"},
	{"EXISTS", "key", "KV"},
	{"EXPIRE", "key seconds", "KV"},
	{"EXPIREAT", "key timestamp", "KV"},
	{"FLUSHALL", "-", "Server"},
	{"FLUSHDB", "-", "Server"},
	{"FULLSYNC", "[NEW]", "Replication"},
	{"GET", "key", "KV"},
	{"GETBIT", "key offset", "KV"},
	{"GETRANGE", "key start end", "KV"},
	{"GETSET", " key value", "KV"},
	{"HCLEAR", "key", "Hash"},
	{"HDEL", "key field [field ...]", "Hash"},
	{"HDUMP", "key", "Hash"},
	{"HEXISTS", "key field", "Hash"},
	{"HEXPIRE", "key seconds", "Hash"},
	{"HEXPIREAT", "key timestamp", "Hash"},
	{"HGET", "key field", "Hash"},
	{"HGETALL", "key", "Hash"},
	{"HINCRBY", "key field increment", "Hash"},
	{"HKEYEXISTS", "key", "Hash"},
	{"HKEYS", "key", "Hash"},
	{"HLEN", "key", "Hash"},
	{"HMCLEAR", "key [key ...]", "Hash"},
	{"HMGET", "key field [field ...]", "Hash"},
	{"HMSET", "key field value [field value ...]", "Hash"},
	{"HPERSIST", "key", "Hash"},
	{"HSET", "key field value", "Hash"},
	{"HTTL", "key", "Hash"},
	{"HVALS", "key", "Hash"},
	{"INCR", "key", "KV"},
	{"INCRBY", "key increment", "KV"},
	{"INFO", "[section]", "Server"},
	{"KEYS", "pattern", "KV"},
	{"LCLEAR", "key", "List"},
	{"LDUMP", "key", "List"},
	{"LEXPIRE", "key seconds", "List"},
	{"LEXPIREAT", "key timestamp", "List"},
	{"LINDEX", "key index", "List"},
	{"LKEYEXISTS", "key", "List"},
	{"LLEN", "key", "List"},
	{"LMCLEAR", "key [key ...]", "List"},
	{"LPERSIST", "key", "List"},
	{"LPOP", "key", "List"},
	{"LPUSH", "key value [value ...]", "List"},
	{"LRANGE", "key start stop", "List"},
	{"LTTL", "key", "List"},
	{"MGET", "key [key ...]", "KV"},
	{"MSET", "key value [key value ...]", "KV"},
	{"PERSIST", "key", "KV"},
	{"PING", "-", "Server"},
	{"RESTORE", "key ttl value", "Server"},
	{"ROLE", "-", "Server"},
	{"RPOP", "key", "List"},
	{"RPUSH", "key value [value ...]", "List"},
	{"SADD", "key member [member ...]", "Set"},
	{"SCARD", "key", "Set"},
	{"SCLEAR", "key", "Set"},
	{"SCRIPT EXISTS", "script [script ...]", "Script"},
	{"SCRIPT FLUSH", "-", "Script"},
	{"SCRIPT LOAD", "script", "Script"},
	{"SDIFF", "key [key ...]", "Set"},
	{"SDIFFSTORE", "destination key [key ...]", "Set"},
	{"SDUMP", "key", "Set"},
	{"SELECT", "index", "Server"},
	{"SET", "key value", "KV"},
	{"SETBIT", "key offset value", "KV"},
	{"SETEX", "key seconds value", "KV"},
	{"SETNX", "key value", "KV"},
	{"SETRANGE", "key offset value", "KV"},
	{"SEXPIRE", "key seconds", "Set"},
	{"SEXPIREAT", "key timestamp", "Set"},
	{"SINTER", "key [key ...]", "Set"},
	{"SINTERSTORE", "destination key [key ...]", "Set"},
	{"SISMEMBER", "key member", "Set"},
	{"SKEYEXISTS", "key", "Set"},
	{"SLAVEOF", "host port [RESTART] [READONLY]", "Replication"},
	{"SMCLEAR", "key [key ...]", "Set"},
	{"SMEMBERS", "key", "Set"},
	{"SPERSIST", "key", "Set"},
	{"SREM", "key member [member ...]", "Set"},
	{"STRLEN", "key", "KV"},
	{"STTL", "key", "Set"},
	{"SUNION", "key [key ...]", "Set"},
	{"SUNIONSTORE", "destination key [key ...]", "Set"},
	{"SYNC", "logid", "Replication"},
	{"TIME", "-", "Server"},
	{"TTL", "key", "KV"},
	{"XHSCAN", "key cursor [MATCH match] [COUNT count] [ASC|DESC]", "Hash"},
	{"XLSORT", "key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC|DESC] [ALPHA] [STORE destination]", "List"},
	{"XSCAN", "type cursor [MATCH match] [COUNT count] [ASC|DESC]", "Server"},
	{"XSSCAN", "key cursor [MATCH match] [COUNT count] [ASC|DESC]", "Set"},
	{"XSSORT", "key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC|DESC] [ALPHA] [STORE destination]", "Set"},
	{"XZSCAN", "key cursor [MATCH match] [COUNT count] [ASC|DESC]", "ZSet"},
	{"XZSORT", "key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC|DESC] [ALPHA] [STORE destination]", "ZSet"},
	{"ZADD", "key score member [score member ...]", "ZSet"},
	{"ZCARD", "key", "ZSet"},
	{"ZCLEAR", "key", "ZSet"},
	{"ZCOUNT", "key min max", "ZSet"},
	{"ZDUMP", "key", "ZSet"},
	{"ZEXPIRE", "key seconds", "ZSet"},
	{"ZEXPIREAT", "key timestamp", "ZSet"},
	{"ZINCRBY", "key increment member", "ZSet"},
	{"ZINTERSTORE", "destkey numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]", "ZSet"},
	{"ZKEYEXISTS", "key", "ZSet"},
	{"ZLEXCOUNT", "key min max", "ZSet"},
	{"ZMCLEAR", "key [key ...]", "ZSet"},
	{"ZPERSIST", "key", "ZSet"},
	{"ZRANGE", "key start stop [WITHSCORES]", "ZSet"},
	{"ZRANGEBYLEX", "key min max [LIMIT offset count]", "ZSet"},
	{"ZRANGEBYSCORE", "key min max [WITHSCORES] [LIMIT offset count]", "ZSet"},
	{"ZRANK", "key member", "ZSet"},
	{"ZREM", "key member [member ...]", "ZSet"},
	{"ZREMRANGBYLEX", "key min max", "ZSet"},
	{"ZREMRANGEBYRANK", "key start stop", "ZSet"},
	{"ZREMRANGEBYSCORE", "key min max", "ZSet"},
	{"ZREVRANGE", "key start stop [WITHSCORES]", "ZSet"},
	{"ZREVRANGEBYSCORE", "key max min  [WITHSCORES][LIMIT offset count]", "ZSet"},
	{"ZREVRANK", "key member", "ZSet"},
	{"ZSCORE", "key member", "ZSet"},
	{"ZTTL", "key", "ZSet"},
	{"ZUNIONSTORE", "destkey numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]", "ZSet"},
}

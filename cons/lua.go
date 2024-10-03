package cons

// 释放锁 Lua 脚本，防止任何客户端都能解锁
const ReleaseLockLuaScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
`

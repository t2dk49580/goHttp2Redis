local a = redis.call('get','a')
local b = redis.call('get','b')
if a == nil then
    redis.call('set','a','1')
    a = redis.call('get','a')
end
if b == nil then
    redis.call('set','b','2')
    b = redis.call('get','b')
end
return a+b
--[[
{
	"id":"xx",//user id
	"data":{
		"msg":"xx"
	}
}
]]--
function _check(pData,pJson)
    local k,v
    for k,v in pairs(pData) do
        if pJson[v] == nil then
            return false
        end
    end
    return true
end
if ARGV[1] == nil then
	return '{"result":"create argv[2] == nil"}'
end
local check = {'id','data'}
local input = cjson.decode(ARGV[1])
if _check(check,input) == false then
	return '{"result":"create check input fail"}'
end

local user = input.id
local msgtype = input.type
local data = input.data

local thesisid = '0'
if redis.call('EXISTS','thesis') > 0 then
    thesisid = redis.call('get','thesis')
end

data.id = user
data.thesis = thesisid
redis.call('rpush',user..'-'..thesisid,cjson.encode(data))
redis.call('rpush',user,thesisid)
redis.call('set','thesis',thesisid+1)
return '{"result":"create ok"}'

--[[
{
	"id":"xx",//user id
	"data":{
		"id":"xx",//reply user id
		"thesis":"xx",//reply thesis id
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
	return '{"result":"reply argv[2] == nil"}'
end
local check = {'id','data'}
local input = cjson.decode(ARGV[1])
if _check(check,input) == false then
	return '{"result":"reply check input fail"}'
end

local user = input.id
local msgtype = input.type
local data = input.data

check = {'id','thesis'}
if _check(check,data) == false then
	return '{"result":"reply check data fail"}'
end

local replyid = data.id
local thesisid = data.thesis

if redis.call('EXISTS',replyid..'-'..thesisid) == 0 then
    return '{"result":"reply user or thesis not found"}'
end

data.id = user
redis.call('rpush',replyid..'-'..thesisid,cjson.encode(data))

data.result = 'reply ok'

return cjson.encode(data)

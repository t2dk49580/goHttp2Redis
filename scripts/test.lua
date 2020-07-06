--[[
{
	"id":"xx",//user id
	"type":"xx",//creat,reply
	"data":{
		"id":"xx",//reply user id
		"thesis":"xx",//reply thesis id
		"msg":"xx"
	}
}
]]--
local input = cjson.decode(ARGV[2])
local user = input.id
local msgtype = input.type
local data = input.data

if msgtype ~= 'create' then
	return '{"result":"message type error"}'
end
local thesisid = redis.call('get','thesis')
data.thesis = thesisid
redis.call('rpush',user..'-'..thesisid,cjson.encode(data))
redis.call('rpush',user,thesisid)
redis.call('set','thesis',thesisid+1)

if msgtype == 'create' then
	local thesisid = redis.call('get','thesis')
	data.thesis = thesisid
	redis.call('rpush',user..'-'..thesisid,cjson.encode(data))
	redis.call('rpush',user,thesisid)
	redis.call('set','thesis',thesisid+1)
elseif msgtype == 'reply' then
	local replyid = data.id
	local thesisid = data.thesis
	redis.call('rpush',replyid..'-'..thesisid,cjson.encode(data.msg))
end

return '{"result":"create ok"}'


local headerToExtract = "foo"

local function extractAndLogHeader(request)
    local headerValue = request:header(headerToExtract)

    if headerValue then
        print("Extracted header value: " .. headerValue)
    else
        print("Header not found: " .. headerToExtract)
    end
end

local hook = {
    on_request = function (self, request)
        extractAndLogHeader(request)
    end,
}

-- Register the Lua hook
cilium.register_lua_event(hook)


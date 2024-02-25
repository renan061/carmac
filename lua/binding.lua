local rollup = {}
rollup.__index = rollup

-- #include "rollup.h"
-- #include "io.h"

function rollup.new()
    local t = {}
    setmetatable(t, rollup)
    -- t.inner = *C.cmt_rollup_t
    return t
end

function rollup:destroy()
    -- cmt_rollup_fini
end

function rollup:finish(accept)
    -- cmt_rollup_finish
end

function rollup:read_advance_state()
    -- cmt_rollup_read_advance_state
end

function rollup:read_inspect_state()
    -- cmt_rollup_read_inspect_state
end

function rollup:emit_voucher()
    -- cmt_rollup_emit_voucher
end

function rollup:emit_notice()
    -- cmt_rollup_emit_notice
end

function rollup:emit_report()
    -- cmt_rollup_emit_report
end

return rollup


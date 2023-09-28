wrk.method = "POST"
wrk.body   = '{"text":"for(var i=0;i<10;i=i+1){print i;}"}'
wrk.headers["Content-Type"] = "application/json"

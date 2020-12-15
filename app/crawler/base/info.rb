class Base::Info



    def get_stock_value
        r = RestClient::Request.execute(
            method: :get, 
            url: 'https://danjuanapp.com/djapi/fundx/activity/user/vip_valuation/show/detail?source=jiucai',
            "headers"=>{
                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.11 Safari/537.36"
            }
        )
        result = JSON.parse(r.body)
        result
    end
end
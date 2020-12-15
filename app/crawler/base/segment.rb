class Base::Segment
    
    def get_segment
        r = RestClient::Request.execute(
            method: :get, 
            url: 'http://quote.eastmoney.com/center/api/sidemenu.json',
            "headers"=>{
                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.11 Safari/537.36"
            }
        )
    end
    



end
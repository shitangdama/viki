class Base::Segment
    
    def get_df_segment
        r = RestClient::Request.execute(
            method: :get, 
            url: 'http://quote.eastmoney.com/center/api/sidemenu.json',
            "headers"=>{
                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.11 Safari/537.36"
            }
        )
    end
    
    def get_sina_segment
        r = RestClient::Request.execute(
            method: :get, 
            url: 'http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodes',
            "headers"=>{
                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.11 Safari/537.36"
            }
        )

    end
end
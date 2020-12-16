class Crawler::Segment
    def get_or_update_segment

        segment = Base::Segment.new()
        result = segment.get_df_segment()

        data =  JSON.parse(result.body)
        all_data = data[5]["next"]
        source = "df"
        
        all_data.each do |board|
            board["next"].each do |item|
                segment_type = board["title"]
                segment = Segment.find_by(
                    key: item["key"]
                )
                if segment.blank?
                    segment = Segment.new
                    segment.segment_type = segment_type
                    segment.key = item["key"]
                    segment.title = item["title"]
                    segment.source = source
                    segment.save()
                end
            end
        end
    end

    def get_sina_segment
        segment = Base::Segment.new()
        result = segment.get_sina_segment()
        source = "sina"

        data =  JSON.parse(result.body)
        sina_data = data[1][0][1][0][1]
        sina_type = data[1][0][1][0][0]


        sw_data = data[1][0][1][1][1]
        sw_type = data[1][0][1][1][0]
        p sw_data

        # sina_data.each do |item|
        #     segment = Segment.find_by(
        #         key: item[2]
        #     )
        #     if segment.blank?
        #         segment = Segment.new
        #         segment.segment_type = sina_type
        #         segment.key = item[2]
        #         segment.title = item[0]
        #         segment.source = source
        #         segment.save()
        #     end
        # end

        
        # sw_data.each do |item|
        #     segment = Segment.find_by(
        #         key: item[2]
        #     )
        #     if segment.blank?
        #         segment = Segment.new
        #         segment.segment_type = sina_type
        #         segment.key = item[2]
        #         segment.title = item[0]
        #         segment.source = source
        #         segment.save()
        #     end
        # end

    end
end
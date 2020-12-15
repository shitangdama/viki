class Crawler::Segment
    def get_or_update_segment

        segment = Base::Segment.new()
        result = segment.get_segment()

        data =  JSON.parse(result.body)
        all_data = data[5]["next"]
        source = "东方财富"
        
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
end
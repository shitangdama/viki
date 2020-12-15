


namespace :test do
    desc "This task does nothing"
    task :generate => :environment do
        info = Base::Info.new()
        result = info.get_stock_value()

        if result.present?
            spread_trends = result["data"]["spread_trends"]

            spread_trends.each do |trend|
                p trend
                time = DateTime.parse(trend["time"]).to_i
                t = SpreadTrend.find_by(time: time)
                if t.blank?
                    t = SpreadTrend.new
                    t.time = time
                    t.spread = trend["spread"]
                    t.wind_all_ashares_range = trend["wind_all_ashares_range"]
                    t.save()
                end
            end

            current_time = DateTime.parse(result["data"]["time"]).to_i
            current_trend = SpreadTrend.find_by(time: current_time)
            p current_trend
            if current_trend.present?
                p current_time

                current_trend.time = current_time
                current_trend.comment = DateTime.parse(result["data"]["comment"])
                current_trend.spread_td = DateTime.parse(result["data"]["spread_td"])
                current_trend.ashares_total_percent = DateTime.parse(result["data"]["ashares_total_percent"])
            end
        end
    end

    desc "Testing environment and variables"
    # rake "hello[World]" 
    task :hello, [:message]  => :environment  do |t, args|
        puts args
    end

    task :shares => :environment do


    end
# RAILS_ENV=development  rake test:segments
    task :segments => :environment do
        Rails.env = "development"
        segment = Crawler::Segment.new()
        segment.get_or_update_segment()
    end
end

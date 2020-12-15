class CreateSpreadTrends < ActiveRecord::Migration[6.0]
  def change
    create_table :spread_trends do |t|
      t.integer    :time,                               null: false
      t.string :comment,                            default: ""
      t.string :spread_td,                          default: ""
      t.float  :spread,                             default: 0.0
      t.float  :ashares_total_percent,              default: 0.0
      t.float  :wind_all_ashares_range,             default: 0.0

      t.timestamps
    end
    add_index :spread_trends, :time,                   unique: true
  end
end

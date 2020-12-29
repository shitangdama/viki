class CreateShares < ActiveRecord::Migration[6.0]
  def change
    create_table :shares do |t|
      # 股票
      t.string :code,  default: "", null: false
      t.string :name,  default: "", null: false
      t.jsonb  :info,  default: {}
      t.timestamps
    end
  end
end

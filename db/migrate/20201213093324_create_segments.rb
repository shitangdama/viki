class CreateSegments < ActiveRecord::Migration[6.0]
  def change
    create_table :segments do |t|
      t.string :source,        default: "", null: false
      t.string :segment_type,  default: "", null: false
      t.string :key,           default: "", null: false
      t.string :title,         default: "", null: false
      t.string :secondary_title, default: ""
      t.integer :shares, default: [], array: true
      t.jsonb  :info,          default: {}
      t.timestamps
    end
  end
end

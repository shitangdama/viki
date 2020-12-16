class CreateSegments < ActiveRecord::Migration[6.0]
  def change
    create_table :segments do |t|
      t.references :parent

      t.string :source,        default: "", null: false
      t.string :segment_type,  default: "", null: false
      t.string :key,           default: "", null: false
      t.string :title,         default: "", null: false
      t.jsonb  :info,          default: {}
      t.timestamps
    end
  end
end

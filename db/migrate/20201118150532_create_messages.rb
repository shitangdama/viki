class CreateMessages < ActiveRecord::Migration[6.0]
  def change
    create_table :messages do |t|
      t.string :title, null: false, default: ""
      t.string :content, default: ""
      t.string :source, default: ""
      t.timestamps
    end
  end
end

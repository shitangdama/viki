# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `rails
# db:schema:load`. When creating a new database, `rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2020_12_13_093324) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "messages", force: :cascade do |t|
    t.string "title", default: "", null: false
    t.string "content", default: ""
    t.string "source", default: ""
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

  create_table "segments", force: :cascade do |t|
    t.string "source", default: "", null: false
    t.string "segment_type", default: "", null: false
    t.string "key", default: "", null: false
    t.string "title", default: "", null: false
    t.string "secondary_title", default: ""
    t.integer "shares", default: [], array: true
    t.jsonb "info", default: {}
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

  create_table "shares", force: :cascade do |t|
    t.string "code", default: "", null: false
    t.string "name", default: "", null: false
    t.integer "concept", default: [], array: true
    t.integer "region", default: [], array: true
    t.integer "industry", default: [], array: true
    t.integer "secondary_industry", default: [], array: true
    t.jsonb "info", default: {}
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

  create_table "spread_trends", force: :cascade do |t|
    t.integer "time", null: false
    t.string "comment", default: ""
    t.string "spread_td", default: "0.0"
    t.float "spread", default: 0.0
    t.float "ashares_total_percent", default: 0.0
    t.float "wind_all_ashares_range", default: 0.0
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["time"], name: "index_spread_trends_on_time", unique: true
  end

  create_table "users", force: :cascade do |t|
    t.string "username", default: "", null: false
    t.string "email", default: "", null: false
    t.string "encrypted_password", default: "", null: false
    t.string "reset_password_token"
    t.datetime "reset_password_sent_at"
    t.datetime "remember_created_at"
    t.boolean "master", default: false
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["email"], name: "index_users_on_email", unique: true
    t.index ["reset_password_token"], name: "index_users_on_reset_password_token", unique: true
  end

end

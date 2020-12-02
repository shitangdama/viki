require "application_system_test_case"

class SpreadTrendsTest < ApplicationSystemTestCase
  setup do
    @spread_trend = spread_trends(:one)
  end

  test "visiting the index" do
    visit spread_trends_url
    assert_selector "h1", text: "Spread Trends"
  end

  test "creating a Spread trend" do
    visit spread_trends_url
    click_on "New Spread Trend"

    click_on "Create Spread trend"

    assert_text "Spread trend was successfully created"
    click_on "Back"
  end

  test "updating a Spread trend" do
    visit spread_trends_url
    click_on "Edit", match: :first

    click_on "Update Spread trend"

    assert_text "Spread trend was successfully updated"
    click_on "Back"
  end

  test "destroying a Spread trend" do
    visit spread_trends_url
    page.accept_confirm do
      click_on "Destroy", match: :first
    end

    assert_text "Spread trend was successfully destroyed"
  end
end

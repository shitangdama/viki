require 'test_helper'

class SpreadTrendsControllerTest < ActionDispatch::IntegrationTest
  setup do
    @spread_trend = spread_trends(:one)
  end

  test "should get index" do
    get spread_trends_url
    assert_response :success
  end

  test "should get new" do
    get new_spread_trend_url
    assert_response :success
  end

  test "should create spread_trend" do
    assert_difference('SpreadTrend.count') do
      post spread_trends_url, params: { spread_trend: {  } }
    end

    assert_redirected_to spread_trend_url(SpreadTrend.last)
  end

  test "should show spread_trend" do
    get spread_trend_url(@spread_trend)
    assert_response :success
  end

  test "should get edit" do
    get edit_spread_trend_url(@spread_trend)
    assert_response :success
  end

  test "should update spread_trend" do
    patch spread_trend_url(@spread_trend), params: { spread_trend: {  } }
    assert_redirected_to spread_trend_url(@spread_trend)
  end

  test "should destroy spread_trend" do
    assert_difference('SpreadTrend.count', -1) do
      delete spread_trend_url(@spread_trend)
    end

    assert_redirected_to spread_trends_url
  end
end

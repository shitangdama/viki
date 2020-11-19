  
class Admin::BaseController < ApplicationController
  before_action :authenticate_user!
  before_action :check_authorization!
  layout 'admin'

  private
  def check_authorization!
    # if current_user
    # if !current_user.check_authorization(controller_name.to_sym, action_name.to_sym)
      # url = request.referer
      #         redirect_to url
      #     end
      # else
      #     url = request.referer || new_user_session_path
      #     redirect_to url
      # end
  end
end
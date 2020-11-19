class AuthenticationController < ActionController::API
  # skip_before_action :configure_permitted_parameters
  def authenticate
    user = User.find_by(name: params[:username])
    user = User.find_for_authentication(name: params[:username])
    if user.present? && user.valid_password?(params[:password])
      render json: user, status: :ok
    else
      render json: { error: "user is not exist" }, status: :unauthorized
    end
  end

  def authed_user
    username = params[:username]
    @user = User.find_by(name: username)
    if @user
      render json: @user, status: 200
    else
      render json: { error: "Resource not found." }, status: 404
    end
  end
end
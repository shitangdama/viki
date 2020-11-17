class AuthenticationController < ActionController::API
    # skip_before_action :configure_permitted_parameters

  before_action :configure_permitted_parameters, if: :devise_controller?

  def configure_permitted_parameters
    update_attrs = %i[password password_confirmation current_password]
    devise_parameter_sanitizer.permit :account_update, keys: update_attrs

    if controller_name == 'registrations'
      unless %w[edit update].include? action_name
        render nothing: true, status: 404
        return
      end
    end
  end

  private
  def after_sign_in_path_for(resource)
    admin_root_path
  end

  def after_sign_out_path_for(resource)
    new_user_session_path
  end
end
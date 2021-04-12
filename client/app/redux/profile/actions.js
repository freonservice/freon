import {
    PROFILE_REQUEST,
    UPDATE_USER_PASSWORD_REQUEST,
    UPDATE_PROFILE_REQUEST,
} from './constants';

export const profileRequest = () => ({
    type: PROFILE_REQUEST,
});

export const updateUserProfileRequest = (email, first_name, second_name, role = 'translator', status = 'active', user_id = 0) => ({
    type: UPDATE_PROFILE_REQUEST,
    user_id: user_id,
    role: role,
    status: status,
    email: email,
    first_name: first_name,
    second_name: second_name,
});

export const updateUserPasswordRequest = (old_password, new_password, repeat_password) => ({
    type: UPDATE_USER_PASSWORD_REQUEST,
    old_password: old_password,
    new_password: new_password,
    repeat_password: repeat_password,
});

import {
    LIST_SETTINGS_REQUEST,
    UPDATE_SETTING_STORAGE_REQUEST,
    UPDATE_SETTING_TRANSLATION_REQUEST,
} from './constants';

export const listSettingsRequest = () => ({
    type: LIST_SETTINGS_REQUEST,
});

export const updateSettingStorageRequest = (use = 0) => ({
    type: UPDATE_SETTING_STORAGE_REQUEST,
    use: use,
});

export const updateTranslationStorageRequest = (use = 0, auto = false, main_language = "en") => ({
    type: UPDATE_SETTING_TRANSLATION_REQUEST,
    use: use,
    auto: auto,
    main_language: main_language,
});

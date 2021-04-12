import {
    LIST_LOCALIZATIONS_REQUEST,
    CREATE_LOCALIZATION_REQUEST,
} from './constants';

export const listLocalizationsRequest = () => ({
    type: LIST_LOCALIZATIONS_REQUEST,
});

export const createLocalizationRequest = (locale, langName) => ({
    type: CREATE_LOCALIZATION_REQUEST,
    locale: locale,
    langName: langName,
});

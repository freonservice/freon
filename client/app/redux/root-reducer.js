import {combineReducers} from 'redux';
import login from './login/reducer';
import profile from './profile/reducer';
import localizations from './localizations/reducer';
import identifiers from './identifiers/reducer';
import categories from './categories/reducer';
import translations from './translations/reducer';
import users from './users/reducer';
import stat from './stat/reducer';
import translationFiles from './translationFiles/reducer';
import languages from './languages/reducer';
import settings from './settings/reducer';

export default combineReducers({
    login,
    profile,
    localizations,
    identifiers,
    categories,
    translations,
    users,
    stat,
    translationFiles,
    languages,
    settings
});

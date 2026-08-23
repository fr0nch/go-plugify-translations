#include "shared.h"

PLUGIFY_EXPORT String (*__translations_LoadTranslation)(Vector*) = NULL;


PLUGIFY_EXPORT void (*__translations_ReloadTransalations)() = NULL;


PLUGIFY_EXPORT bool (*__translations_HasTranslation)(String*) = NULL;


PLUGIFY_EXPORT bool (*__translations_HasTranslatedForLanguage)(String*, String*) = NULL;


PLUGIFY_EXPORT String (*__translations_Translate)(String*, String*) = NULL;


PLUGIFY_EXPORT String (*__translations_TranslateFormat)(String*, String*, Vector*) = NULL;



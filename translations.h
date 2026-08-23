#pragma once

#include "shared.h"

extern String (*__translations_LoadTranslation)(Vector*);

static String LoadTranslation(Vector* paths) {
	return __translations_LoadTranslation(paths);
}

extern void (*__translations_ReloadTransalations)();

static void ReloadTransalations() {
	__translations_ReloadTransalations();
}

extern bool (*__translations_HasTranslation)(String*);

static bool HasTranslation(String* key) {
	return __translations_HasTranslation(key);
}

extern bool (*__translations_HasTranslatedForLanguage)(String*, String*);

static bool HasTranslatedForLanguage(String* key, String* lang) {
	return __translations_HasTranslatedForLanguage(key, lang);
}

extern String (*__translations_Translate)(String*, String*);

static String Translate(String* lang, String* key) {
	return __translations_Translate(lang, key);
}

extern String (*__translations_TranslateFormat)(String*, String*, Vector*);

static String TranslateFormat(String* lang, String* key, Vector* args) {
	return __translations_TranslateFormat(lang, key, args);
}


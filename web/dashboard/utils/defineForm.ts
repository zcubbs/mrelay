import * as z from 'zod'
import { useForm } from 'vee-validate'
import i18next from 'i18next'
import { zodI18nMap } from 'zod-i18n-map'
import { toTypedSchema } from '@vee-validate/zod'
import translationFr from 'zod-i18n-map/locales/en/zod.json'
import translationEn from 'zod-i18n-map/locales/fr/zod.json'

/**
 * This function is used to define a form
 */
export default () => {
  const { locale } = useI18n()

  // This function creates a form based on the provided schema.
  const createForm = (schema: Record<any, any>) => {
    // Convert the schema to a typed schema using z.object.
    const formSchema = toTypedSchema(z.object(schema))

    // Return a form instance with validation schema.
    return useForm({ validationSchema: formSchema })
  }

  // This function initializes zod with the current locale
  const initZodLocal = (lang: string) => {
    // Initialize i18next with the specified locale and translation resources.
    i18next.init({
      lng: lang,
      resources: {
        en: { zod: translationFr }, // English translations
        fr: { zod: translationEn }, // French translations
      },
    })

    // Set the error map using zodI18nMap.
    z.setErrorMap(zodI18nMap)

    return z
  }

  return { z: initZodLocal(locale.value), createForm, updateFormLocal: initZodLocal }
}

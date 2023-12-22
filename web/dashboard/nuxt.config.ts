export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/i18n', '@vueuse/nuxt', 'nuxt-phosphor-icons'],
  i18n: {
    lazy: true,
    langDir: 'lang',
    defaultLocale: 'en',
    detectBrowserLanguage: {
      useCookie: true,
    },
    strategy: 'no_prefix',
    locales: [
      {
        code: 'en',
        files: ['en/messages.json'],
      },
      {
        code: 'fr',
        files: ['fr/messages.json'],
      },
    ],
  },
  components: [
    {
      path: '~/components/shad',
      // this is required else Nuxt will autoImport `.ts` file
      extensions: ['.vue'],
      // prefix for your components, eg: shad-button
      prefix: 'Shad',
    },
    {
      path: '~/components/ui',
      // this is required else Nuxt will autoImport `.ts` file
      extensions: ['.vue'],
      // prefix for your components, eg: ui-button
      prefix: 'Ui',
    },
  ],
})

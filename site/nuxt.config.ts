// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  srcDir: 'src/',

  ssr: false,
  modules: [
    '@pinia/nuxt',
    '@vueuse/nuxt',
    // https://color-mode.nuxtjs.org/#configuration
    '@nuxtjs/color-mode',
    '@element-plus/nuxt',
    ['nuxt-lazy-load', {
      images: true,
      videos: true,
      audios: true,
      iframes: true,
      native: true,
      directiveOnly: false,

      // Default image must be in the public folder
      // defaultImage: '/images/default-image.jpg',

      // To remove class set value to false
      loadingClass: 'isLoading',
      loadedClass: 'isLoaded',
      appendClass: 'lazyLoad',

      observerConfig: {
        // See IntersectionObserver documentation
      },
      
    }],
  ],
  build: {
    transpile: ['pinia'],
  },
  plugins: [
  ],

  elementPlus: {
    defaultLocale: 'zh-cn',
  },

  colorMode: {
    preference: 'system', // default value of $colorMode.preference
    fallback: 'light', // fallback value if not system preference found
    storageKey: 'duck-color-mode',
    classPrefix: 'theme-',
    classSuffix: '',
  },

  imports: {
    dirs: [
      'apis',
      'stores',
    ],
  },

  app: {
    head: {
      title: 'duck',
      htmlAttrs: { class: 'theme-light has-navbar-fixed-top' },
      // script: [
      //   {
      //     src: 'https://hm.baidu.com/hm.js?f14b836e09b72aedce29a86e809936de',
      //     type: 'text/javascript',
      //     async: true
      //   }
      // ]
    },
  },
  css: [
    '~/assets/css/index.scss',
  ],

  nitro: {
    routeRules: {
      '/api/**': {
        proxy: `${import.meta.env.SERVER_URL}/api/**`,
      },
      '/admin/**': {
        proxy: `${import.meta.env.SERVER_URL}/admin/**`,
      },
    },
  },

  compatibilityDate: '2025-03-15',
})
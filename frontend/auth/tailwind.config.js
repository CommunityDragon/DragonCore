/*
** TailwindCSS Configuration File
**
** Docs: https://tailwindcss.com/docs/configuration
** Default: https://github.com/tailwindcss/tailwindcss/blob/master/stubs/defaultConfig.stub.js
*/
module.exports = {
  theme: {
    extend: {
      screens: {
        'dark': { raw: '(prefers-color-scheme: dark)' }
      },
      spacing: {
        '72': '18rem',
        '80': '20rem',
        '88': '22rem',
        '96': '24rem',
        '104': '26rem',
        '112': '28rem',
        '120': '30rem',
        '128': '32rem',
        '136': '34rem',
        '144': '36rem',
        '152': '38rem',
        '160': '40rem',
        '168': '42rem',
        '176': '44rem',
        '184': '46rem',
        '192': '48rem',
        '200': '50rem',
        '208': '52rem',
        '216': '54rem',
        '224': '56rem',
        '232': '58rem',
        '240': '60rem',
        '248': '62rem',
        '256': '64rem',
      },
      maxWidth: theme => ({
        ...theme('spacing'),
      }),
      colors: {
        gray: {
          100: '#F5F5F5',
          200: '#EEEEEE',
          300: '#E0E0E0',
          400: '#BDBDBD',
          500: '#9E9E9E',
          600: '#757575',
          700: '#616161',
          800: '#424242',
          900: '#212121',
          950: '#151515',
        },
      },
      placeholderColor: theme => theme('colors'),
    },
  },
}
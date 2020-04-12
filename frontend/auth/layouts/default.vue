<template>
  <div class="h-full">
    <div class="bg-animated"></div>
    <div class="bg-gradient"></div>
    <nuxt />
  </div>
</template>

<style lang="postcss">
html {
  @apply antialiased leading-snug;
  @apply bg-gray-400;

  @screen dark {
    @apply bg-gray-950;
  }
}

body {
  @apply flex flex-col min-h-screen;
}

#__nuxt,
#__layout {
  @apply h-screen;
}
</style>
<style lang="postcss" scoped>
@keyframes move-bg {
  from {
    background-position: 0 0;
  }
  to {
    background-position: 1200px 600px;
  }
}

.bg-animated,
.bg-gradient {
  @apply fixed top-0 left-0 w-full h-full;
  z-index: -1;
}

.bg-gradient {
  background: linear-gradient(to right, #bdbdbd 50%, transparent 50%);

  @screen dark {
    background: linear-gradient(to right, #212121 40%, transparent);
  }
}

.bg-animated {
  @apply bg-gray-900;
  animation: move-bg 60s infinite;
  animation-timing-function: linear;
  background-size: 600px;
  background-image: url('~assets/img/bg-light.svg');

  @screen dark {
    @apply bg-black;
    background-image: url('~assets/img/bg-dark.svg');
  }
}
</style>

<template>
  <div class="container">
    <div class="p-8 inline-block flex w-full justify-center">
      <card width="160">
        <div class="card__content">
          <div class="card__half flex justify-center">
            <img
              src="~assets/img/logo-white.png"
              width="200"
              class="object-contain"
            />
          </div>
          <div class="card__half">
            <h1 class="card__title">Welcome Back</h1>
            <br />
            <form>
              <form-input label="Username/Email" />
              <form-input label="Password" type="password" />
              <a href="#">Forgot your password?</a>
              <br />
              <button
                href="#"
                class="bg-gray-900 text-gray-100 px-5 py-3 w-full block font-semibold rounded"
              >
                Sign In
              </button>
              Need an account? <a href="#">Register</a>
            </form>
          </div>
        </div>
      </card>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import { FormInput } from '~/components/forms'
import { Card } from '~/components/cards'

@Component({
  components: { FormInput, Card }
})
export default class HomePage extends Vue {}
</script>

<style lang="postcss" scoped>
.container {
  @apply flex items-center justify-center h-full m-auto;
}

.card__title {
  @apply text-center text-3xl;
  font-family: 'IBM Plex Sans', sans-serif;
}

.card__content {
  @apply flex;
  @apply w-full;

  > * {
    @apply p-4;
  }

  .card__half {
    @apply w-1/2;
  }
}
</style>

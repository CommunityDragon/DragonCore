<template>
  <div
    class="card"
    :class="{
      'card--fluid': fluid
    }"
  >
    <slot name="body">
      <div class="card__body">
        <slot></slot>
      </div>
    </slot>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component({
  props: ['fluid', 'width']
})
export default class Card extends Vue {
  /**
   * gives the card a fluid width
   */
  fluid!: boolean

  /**
   * the width of the card
   */
  width?: number | string
}
</script>

<style lang="postcss" scoped>
.card {
  @apply text-left flex bg-white shadow-md text-black rounded-lg overflow-hidden w-full max-w-160;
  @screen dark {
    @apply bg-gray-950 text-gray-100;
  }

  &.card--fluid {
    @apply max-w-full;
  }

  .card__body {
    @apply p-4 w-full;
  }
}
</style>

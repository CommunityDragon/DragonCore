<template>
  <fragment>
    <span class="group">
      <label v-if="label">{{ label }}</label>
      <input
        v-bind="{ ...$props, ...$attrs }"
        :type="$attrs.type ? $attrs.type : 'text'"
      />
      <br />
    </span>
  </fragment>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component({
  props: ['label']
})
export default class FormInput extends Vue {
  /**
   * Form Label
   */
  label?: string
}
</script>

<style lang="postcss" scoped>
label {
  @apply px-1 text-sm text-gray-400;
  font-family: 'Jetbrains Mono', monospace;
}

input {
  @apply block px-3 pb-2 pt-1 w-full bg-transparent border-b-2 border-gray-600;
  font-family: 'IBM Plex Sans', sans-serif;

  &::placeholder {
    @apply text-gray-600;
  }

  &:focus {
    @apply border-gray-400 outline-none;
  }
}
</style>

<template>
  <ValidationProvider :vid="vid" :name="$attrs.label" :rules="rules">
      <div class="block" slot-scope="{ errors }">
        <slot />
        <p class="has-text-danger">{{ errors[0] }}</p>
      </div>
  </ValidationProvider>
</template>

<script>
import { ValidationProvider } from "vee-validate";

export default {
  components: {
    ValidationProvider
  },
  props: {
    vid: {
      type: String
    },
    rules: {
      type: [Object, String],
      default: ''
    }
  }
};
</script>

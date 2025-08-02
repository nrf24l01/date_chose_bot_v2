<template>
  <div class="tz-demo-wrap">
    <div class="dp-container-wrap">
      <VueDatePicker 
        :model-value="modelValue"
        @update:model-value="onUpdate"
        :dark="false"
        :timezone="'Europe/Moscow'" 
        inline 
        auto-apply
        :max-date="props.maxDate"
        :min-date="props.minDate"
        :enable-time-picker="false"
        :disabled="props.disabled"
        multi-dates
      />
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'

const props = defineProps({
  minDate: {
    type: String,
    default: '2023-01-01T00:00:00Z'
  },
  maxDate: {
    type: String,
    default: '2023-12-31T23:59:59Z'
  },
  modelValue: {
    type: Array,
    default: () => []
  },
  disabled: {
    type: Boolean,
    default: false
  }
});

const emit = defineEmits(['update:modelValue']);

function onUpdate(value) {
  emit('update:modelValue', value);
}
</script>
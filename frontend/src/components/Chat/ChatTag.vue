<script setup>
import { ref } from 'vue';

const props = defineProps({
  tags: {
    type: Array,
    required: true
  },
  selectedTag: {
    type: String,
    required: true
  },
});

const emit = defineEmits(['update:selectedTag']);

const selectTag = (tag) => {
  emit('update:selectedTag', tag);
};

const showMoreTags = () => {
  selectTag('更多');
};
</script>

<template>
  <div class="tags">
    <n-tag
      v-for="tag in tags"
      :key="tag"
      :checked="tag === selectedTag"
      @click="selectTag(tag)"
      checkable
    >
      {{ tag }}
    </n-tag>
    <n-tag 
      @click="showMoreTags"
      :checked="selectedTag === '更多'"
      checkable
    > 
      更多
    </n-tag>
  </div>
</template>

<style scoped>
.tags {
  display: flex;
  gap: 0.5em;
  margin-bottom: 1em;
}

button {
  padding: 0.5em 1em;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button.selected {
  background-color: #42b983;
  color: white;
}

button:hover {
  background-color: #369f6b;
}
</style>
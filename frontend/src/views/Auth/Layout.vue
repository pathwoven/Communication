<template>
  <div class="auth-layout">
    <section>
      <router-view />
    </section>
  </div>
</template>

<script setup>
</script>

<style scoped>
.auth-layout {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f2f5;
}

section {
  background-color: white;
  padding: 2em;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  max-width: 400px;
  width: 100%;
  max-height: 800px;
}

section h1 {
  text-align: center;
  margin-bottom: 1em;
}

</style>
<template>
  <div>
    <div v-if="pending">
      <p>Loading course details...</p>
    </div>
    <div v-else-if="error">
      <p class="text-red-500">Could not load the course.</p>
    </div>
    <div v-else-if="course">
      <h1 class="text-4xl font-bold mb-4">{{ course.Title }}</h1>
      <p class="text-lg text-gray-500 dark:text-gray-400">
        {{ course.Description }}
      </p>
    </div>
  </div>
</template>

<script setup>
const route = useRoute();
const token = useCookie("token");
const courseId = route.params.id;

const {
  data: course,
  pending,
  error,
} = await useFetch(`/api/v1/courses/${courseId}`, {
  baseURL: "http://localhost:8080",
  headers: {
    Authorization: `Bearer ${token.value}`,
  },
});
</script>

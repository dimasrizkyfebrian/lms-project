<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold">Available Courses</h1>
      <UButton to="/courses/create" icon="i-heroicons-plus-circle" size="lg">
        Create Course
      </UButton>
    </div>

    <div v-if="pending">
      <p>Loading courses...</p>
    </div>

    <div v-else-if="error">
      <p class="text-red-500">
        Failed to load courses. Please try again later.
      </p>
    </div>

    <div
      v-else-if="courses"
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"
    >
      <UCard v-for="course in courses" :key="course.ID">
        <template #header>
          <h2 class="text-lg font-semibold">{{ course.Title }}</h2>
        </template>
        <p class="text-gray-500 dark:text-gray-400">{{ course.Description }}</p>
        <template #footer>
          <UButton color="primary" :to="`/courses/${course.ID}`">
            View Course
          </UButton>
        </template>
      </UCard>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from "vue";

const token = useCookie("token");

const {
  data: courses,
  pending,
  error,
  refresh,
} = await useFetch("/api/v1/courses", {
  baseURL: "http://localhost:8080",
  headers: {
    Authorization: `Bearer ${token.value}`,
  },
});

onMounted(() => {
  refresh();
});
</script>

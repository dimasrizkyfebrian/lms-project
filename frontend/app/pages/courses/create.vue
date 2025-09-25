<template>
  <div>
    <h1 class="text-3xl font-bold mb-6">Create a New Course</h1>

    <UCard>
      <UForm :state="state" @submit="submit">
        <UFormField label="Course Title" name="title" class="mb-4">
          <UInput
            v-model="state.title"
            placeholder="e.g., Introduction to Go"
          />
        </UFormField>

        <UFormField label="Course Description" name="description">
          <UTextarea
            v-model="state.description"
            placeholder="A brief summary of the course..."
          />
        </UFormField>

        <UButton type="submit" class="mt-6" :loading="isLoading">
          Save Course
        </UButton>
      </UForm>
    </UCard>
  </div>
</template>

<script setup>
const token = useCookie("token");
const toast = useToast();
const router = useRouter();
const isLoading = ref(false);

const state = ref({
  title: "",
  description: "",
});

async function submit() {
  isLoading.value = true;
  try {
    await $fetch("/api/v1/courses", {
      method: "POST",
      baseURL: "http://localhost:8080",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
      body: {
        title: state.value.title,
        description: state.value.description,
      },
    });

    toast.add({ title: "Course created successfully!" });

    // Redirect to the courses page after creation
    router.push("/courses");
  } catch (error) {
    toast.add({
      title: "Error",
      description: error.data?.error || "Failed to create course.",
      color: "red",
    });
  } finally {
    isLoading.value = false;
  }
}
</script>

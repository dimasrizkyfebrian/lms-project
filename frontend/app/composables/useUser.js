// composables/useUser.js
import { useState } from "#app";

export const useUser = () => {
  // useState creates a reactive state that is shared across the application.
  // 'user' is the unique key.
  const user = useState("user", () => null);
  return user;
};

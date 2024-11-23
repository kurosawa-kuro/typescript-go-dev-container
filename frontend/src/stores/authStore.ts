import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';

interface User {
  id: number;
  email: string;
  role: string;
}

interface AuthState {
  user: User | null;
  isLoading: boolean;
  error: string | null;
  setUser: (user: User | null) => void;
  setError: (error: string | null) => void;
  login: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  resetStore: () => void;
}

const initialState = {
  user: null,
  isLoading: false,
  error: null,
};

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      ...initialState,

      setUser: (user) => set({ user }),
      setError: (error) => set({ error }),
      resetStore: () => set(initialState),

      login: async (email: string, password: string) => {
        set({ isLoading: true, error: null });
        try {
          const response = await fetch('http://localhost:8000/api/auth/login', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({ email, password }),
          });

          const data = await response.json();

          if (!response.ok) {
            throw new Error(data.error || 'Login failed');
          }

          set({ user: data.user });
        } catch (error) {
          set({ error: error instanceof Error ? error.message : 'Login failed' });
        } finally {
          set({ isLoading: false });
        }
      },

      logout: async () => {
        try {
          await fetch('http://localhost:8000/api/auth/logout', {
            method: 'POST',
            credentials: 'include',
          });
          clearAuthData();
        } catch (error) {
          console.error('Logout failed:', error);
        }
      },
    }),
    {
      name: 'auth-storage',
      storage: createJSONStorage(() => localStorage),
      partialize: (state) => ({ user: state.user }),
    }
  )
);

function clearAuthData() {
  localStorage.removeItem('auth-storage');
  document.cookie.split(";").forEach((c) => {
    document.cookie = c
      .replace(/^ +/, "")
      .replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
  });
  useAuthStore.persist.clearStorage();
  useAuthStore.getState().resetStore();
}
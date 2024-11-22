'use client';

import { useState } from 'react';
import MicropostModal from '@/app/components/MicropostModal';

export default function CreateMicropostButton() {
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleCreateMicropost = async (title: string) => {
    try {
      const response = await fetch('http://backend:8000/microposts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title }),
      });

      if (!response.ok) {
        throw new Error('Failed to create micropost');
      }

      // TODO: 投稿成功後の処理（例：リストの更新やトースト表示など）
      console.log('Micropost created successfully');
      
      // モーダルを閉じる
      setIsModalOpen(false);
    } catch (error) {
      console.error('Error creating micropost:', error);
      // TODO: エラー処理（例：エラーメッセージの表示など）
    }
  };

  return (
    <>
      <button
        onClick={() => setIsModalOpen(true)}
        className="px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        New Post
      </button>

      <MicropostModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onSubmit={handleCreateMicropost}
      />
    </>
  );
}
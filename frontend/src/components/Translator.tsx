import React, { useState } from "react";
import axios from "axios";

interface TranslateResponse {
  translatedText: string;
}

export const Translator: React.FC = () => {
  const [text, setText] = useState<string>('');
  const [translated, setTranslated] = useState<string>("");

  const handleTranslate = async () => {
    try {
      const res = await axios.post<TranslateResponse>("/api/generate", { text });
      setTranslated(res.data.translatedText);
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <div>
      <textarea
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="翻訳したい文章を入力"
      />
      <button onClick={handleTranslate}>翻訳する</button>
      <div>翻訳結果: {translated}</div>
    </div>
  )
}


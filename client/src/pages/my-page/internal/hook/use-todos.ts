import { useEffect, useState } from "react";
import { serverFetch } from "../../../../utils/fetch";

type Todo = {
    id: number;
    title: string;
    completed: boolean;
};

export const  useTodos = (todoListID: number) => {
    const [todos,setTodos] = useState<Todo[]>([]);
    const [todoListName,setTodoListName] = useState("");

    const fetchTodos = async () => {
        var res = await serverFetch(`/api/todos?todo_list_id=${todoListID}`);
        setTodos(await res.json());
        
        res = await serverFetch(`/api/todo-lists/${todoListID}`);
        const data = await res.json();
        setTodoListName(data.Name);
    };

    useEffect(() => {
        fetchTodos();
    },[todoListID]);

    return {
        todos,
        todoListName,
        fetchTodos,
    };
};
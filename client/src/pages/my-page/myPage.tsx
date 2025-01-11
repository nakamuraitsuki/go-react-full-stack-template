import { useAuth } from "../../provider/auth";
import { TodoList } from "./internal/components/todo-list";
import { TodoCreateForm } from "./internal/components/todo-create-form";
import { useTodos } from "./internal/hook/use-todos";
import { TodoLists } from "./internal/components/todo-lists";

export const MyPage = () => {
    const { user } = useAuth();
    if (!user) return null;

    console.log(user);
    const { todos, fetchTodos } = useTodos(user.defaultTodoListID)
    
    console.log(todos)
    return (
        <div>
            <h1>{user.name}のマイページ</h1>
            <div>
                <h3>{user.name}</h3>
                <TodoLists userID={user.id}/>
            </div>
            <div>
                <TodoCreateForm todoListID={user.defaultTodoListID} refetch={fetchTodos} />
                <TodoList todos={todos} refetch={fetchTodos} />
            </div>
        </div>
    )
}
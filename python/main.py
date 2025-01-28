from langchain import hub
from langchain.agents import AgentExecutor, create_json_chat_agent
from langchain_community.tools.tavily_search import TavilySearchResults
from langchain_core.callbacks.manager import CallbackManager
from langchain_core.callbacks.streaming_stdout import StreamingStdOutCallbackHandler
from langchain_community.chat_models import ChatZhipuAI
from langchain_core.messages import AIMessage, HumanMessage, SystemMessage

tools = [TavilySearchResults(max_results=1)]
prompt = hub.pull("hwchase17/react-chat-json")
llm = ChatZhipuAI(temperature=0.01, model="glm-4")

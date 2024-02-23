// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file llm/service.proto (package llm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ChatCompletionRequest, ChatCompletionResponse, CreateProviderRequest, CreateProviderResponse, GetProviderConfigRequest, GetProviderConfigResponse, GetProviderRequest, GetProviderResponse, LogResponse, ModelRequest, ModelResponse, ProvidersResponse, UpdateProviderRequest, UpdateProviderResponse } from "./service_pb";
import { Empty, MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service llm.v1.LLMService
 */
export const LLMService = {
  typeName: "llm.v1.LLMService",
  methods: {
    /**
     * @generated from rpc llm.v1.LLMService.ChatCompletions
     */
    chatCompletions: {
      name: "ChatCompletions",
      I: ChatCompletionRequest,
      O: ChatCompletionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc llm.v1.LLMService.StreamChatCompletions
     */
    streamChatCompletions: {
      name: "StreamChatCompletions",
      I: ChatCompletionRequest,
      O: ChatCompletionResponse,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * @generated from rpc llm.v1.LLMService.ListModels
     */
    listModels: {
      name: "ListModels",
      I: ModelRequest,
      O: ModelResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc llm.v1.LLMService.ListProviders
     */
    listProviders: {
      name: "ListProviders",
      I: Empty,
      O: ProvidersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc llm.v1.LLMService.GetProvider
     */
    getProvider: {
      name: "GetProvider",
      I: GetProviderRequest,
      O: GetProviderResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc llm.v1.LLMService.CreateProvider
     */
    createProvider: {
      name: "CreateProvider",
      I: CreateProviderRequest,
      O: CreateProviderResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc llm.v1.LLMService.UpsertProvider
     */
    upsertProvider: {
      name: "UpsertProvider",
      I: UpdateProviderRequest,
      O: UpdateProviderResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc llm.v1.LLMService.GetProviderConfig
     */
    getProviderConfig: {
      name: "GetProviderConfig",
      I: GetProviderConfigRequest,
      O: GetProviderConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc llm.v1.LLMService.ListTrackingLogs
     */
    listTrackingLogs: {
      name: "ListTrackingLogs",
      I: Empty,
      O: LogResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;


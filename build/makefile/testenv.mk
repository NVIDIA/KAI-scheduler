
ENVTEST_K8S_VERSION = 1.32.0
ENVTEST_VERSION=release-0.20

envtest-docker-go: gocache
	@ ${ECHO_COMMAND} ${GREEN_CONSOLE} "${CONSOLE_PREFIX} Running unit-tests" ${BASE_CONSOLE}
	${DOCKER_GO_COMMAND} make envtest-go || ${FAILURE_MESSAGE_HANDLER}
	${SUCCESS_MESSAGE_HANDLER}

envtest-go: envtest
	@ ${ECHO_COMMAND} ${GREEN_CONSOLE} "${CONSOLE_PREFIX} Running unit-tests" ${BASE_CONSOLE}
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) -p path --bin-dir $(LOCALBIN))" \
    go test ./... -timeout 30m || ${FAILURE_MESSAGE_HANDLER}
	${SUCCESS_MESSAGE_HANDLER}

ENVTEST = $(LOCALBIN)/setup-envtest
.PHONY: envtest
envtest: ## Download envtest locally if necessary.
	GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-runtime/tools/setup-envtest@${ENVTEST_VERSION}